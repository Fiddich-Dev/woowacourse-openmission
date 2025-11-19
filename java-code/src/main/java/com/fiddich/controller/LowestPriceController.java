package com.fiddich.controller;

import com.fiddich.client.PaymentClient;
import com.fiddich.model.*;
import com.fiddich.model.dto.request.PartitionRequest;
import com.fiddich.model.dto.request.PaymentRequest;
import com.fiddich.model.dto.resonse.DiscountInfoResponse;
import com.fiddich.model.dto.resonse.ExchangeRateResponse;
import com.fiddich.model.dto.resonse.PartitionResponse;
import com.fiddich.client.DiscountPolicyClient;
import com.fiddich.client.ExchangeRateClient;
import com.fiddich.client.PartitionClient;
import com.fiddich.view.InputParser;
import com.fiddich.view.InputView;
import com.fiddich.view.OutputView;

import java.util.List;
import java.util.concurrent.Callable;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

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
        // 쿠폰 할인정보 조회
        List<DiscountInfoResponse> couponDiscountInfoResponseList = DiscountPolicyClient
                .getByCoupon()
                .getContent();
        outputView.printCouponDiscountInfo(couponDiscountInfoResponseList);

        // 카드 할인정보 조회
        List<DiscountInfoResponse> cardDiscountInfoResponseList = DiscountPolicyClient
                .getByCard()
                .getContent();
        outputView.printCouponDiscountInfo(cardDiscountInfoResponseList);

        // 상품 입력 받기
        List<Goods> goodsList = inputGoodsList();

        // 최적의 그룹화
        List<PartitionResponse> partitionResponseList = PartitionClient
                .partitionGoods(new PartitionRequest(goodsList))
                .getContent();
        outputView.printDiscount(partitionResponseList);

        // 환율 조회
        ExchangeRateResponse exchangeRateResponse = ExchangeRateClient
                .getUsdExchangeRate()
                .getContent();
        outputView.printExchangeRate(exchangeRateResponse);

        // 결제
        ExecutorService executorService = Executors.newFixedThreadPool(10); // 스레드 풀 생성
        for (PartitionResponse partitionResponse : partitionResponseList) {
            // 작업 부여
            Callable<ResponseFormat<Void>> task = () -> {
                return PaymentClient.paymentGoods(new PaymentRequest(partitionResponse.goods(), partitionResponse.afterPrice()));
            };
            // 작업 실행
            executorService.submit(task);
        }


        try {
            Thread.sleep(10000);
        } catch (InterruptedException e) {
            System.out.println("asd");
        }
    }

    private List<Goods> inputGoodsList() {
        Long goodsCount = inputParser.stringToLong(inputView.requestGoodsCount());
        return inputView.requestGoods(goodsCount).stream()
                .map(inputParser::toGoods)
                .toList();
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