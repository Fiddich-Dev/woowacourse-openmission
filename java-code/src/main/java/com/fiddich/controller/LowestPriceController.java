package com.fiddich.controller;

import com.fiddich.Console;
import com.fiddich.client.*;
import com.fiddich.model.*;
import com.fiddich.model.dto.request.PartitionRequest;
import com.fiddich.model.dto.resonse.DiscountInfoResponse;
import com.fiddich.model.dto.resonse.ExchangeRateResponse;
import com.fiddich.model.dto.resonse.PartitionResponse;
import com.fiddich.util.ResponseHandler;
import com.fiddich.view.InputParser;
import com.fiddich.view.InputView;
import com.fiddich.view.OutputView;

import java.util.List;

public class LowestPriceController {
    private final InputView inputView;
    private final InputParser inputParser;
    private final OutputView outputView;

    public LowestPriceController(InputView inputView, InputParser inputParser, OutputView outputView) {
        this.inputView = inputView;
        this.inputParser = inputParser;
        this.outputView = outputView;
    }

    public void run() {
        try {
            // 쿠폰 할인정보 조회
            List<DiscountInfoResponse> couponDiscountInfoResponseList = ResponseHandler
                    .handle(DiscountPolicyClient.getByCoupon());
            outputView.printCouponDiscountInfo(couponDiscountInfoResponseList);

            // 카드 할인정보 조회
            List<DiscountInfoResponse> cardDiscountInfoResponseList = ResponseHandler
                    .handle(DiscountPolicyClient.getByCard());
            outputView.printCardDiscountInfo(cardDiscountInfoResponseList);

            // 상품 입력 받기
            List<Goods> goodsList = inputGoodsList();

            // 최적의 그룹화
            List<PartitionResponse> partitionResponseList = ResponseHandler
                    .handle(PartitionClient.partitionGoods(new PartitionRequest(goodsList)));
            outputView.printDiscount(partitionResponseList);

            // 환율 조회
            ExchangeRateResponse exchangeRateResponse = ResponseHandler
                    .handle(ExchangeRateClient.getUsdExchangeRate());
            outputView.printExchangeRate(exchangeRateResponse);

            // 결제 동의시 결제 진행
            if(inputParser.isConfirmed(inputView.requestConfirmPayment())) {
                executePayments(partitionResponseList);
            }
        } finally {
            Console.close();
            System.out.println("프로그램 종료");
        }
    }

    private List<Goods> inputGoodsList() {
        Long goodsCount = inputParser.stringToLong(inputView.requestGoodsCount());
        return inputView.requestGoods(goodsCount).stream()
                .map(inputParser::toGoods)
                .toList();
    }

    private void executePayments(List<PartitionResponse> partitionResponseList) {
        PaymentExecutor executor = new PaymentExecutor(10);

        // 성공한 결제 건만 결과를 받음
        List<PartitionResponse> successList = executor.executePayments(partitionResponseList);

        // 결제 내역 출력
        for (PartitionResponse p : successList) {
            outputView.printPaymentResult(p);
        }
    }
}

/*
A 10
B 20
C 30
D 40
A 10
B 20
C 30
D 40
*/