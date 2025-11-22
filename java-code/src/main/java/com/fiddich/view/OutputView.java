package com.fiddich.view;

import com.fiddich.model.*;
import com.fiddich.model.dto.resonse.DiscountInfoResponse;
import com.fiddich.model.dto.resonse.ExchangeRateResponse;
import com.fiddich.model.dto.resonse.PartitionResponse;

import java.util.List;

public class OutputView {

    private static final String COUPON_DISCOUNT_MESSAGE = "쿠폰 할인 정보";
    private static final String CARD_DISCOUNT_MESSAGE = "카드 할인 정보";
    private static final String JOIN_DELIMITER = ", ";
    private static final String DISCOUNT_FORMAT = "%s달러 이상 구매하면 %s달러 할인%n";
    private static final String DISCOUNT_RESULT_FORMAT = "%s (Before: %s, After: %s)%n";
    private static final String EXCHANGE_RATE_TITLE_MESSAGE = "현재 환율 정보";
    private static final String EXCHANGE_RATE_FORMAT = "1%s = %s원%n";
    private static final String PAYMENT_SUCCESS_FORMAT = "%s : %s달러 결제 성공%n";

    public void printCouponDiscountInfo(List<DiscountInfoResponse> discountInfoResponseList) {
        System.out.println(COUPON_DISCOUNT_MESSAGE);
        printDiscountInfo(discountInfoResponseList);
        System.out.println();
    }

    public void printCardDiscountInfo(List<DiscountInfoResponse> discountInfoResponseList) {
        System.out.println(CARD_DISCOUNT_MESSAGE);
        printDiscountInfo(discountInfoResponseList);
        System.out.println();
    }

    private void printDiscountInfo(List<DiscountInfoResponse> discountInfoResponseList) {
        for (DiscountInfoResponse discountInfoResponse : discountInfoResponseList) {
            String threshold = discountInfoResponse.threshold();
            String discountAmount = discountInfoResponse.discountAmount();
            System.out.printf(DISCOUNT_FORMAT, threshold, discountAmount);
        }
    }

    public void printDiscount(List<PartitionResponse> partitionResponseList) {
        for(PartitionResponse partitionResponse : partitionResponseList) {
            List<String> goodsNames = partitionResponse.goodsList().stream()
                    .map(Goods::getName)
                    .toList();

            String joinedGoodsNames = String.join(JOIN_DELIMITER, goodsNames);
            String beforePrice = partitionResponse.beforePrice();
            String afterPrice = partitionResponse.afterPrice();
            System.out.printf(DISCOUNT_RESULT_FORMAT, joinedGoodsNames, beforePrice, afterPrice);
        }
        System.out.println();
    }

    public void printExchangeRate(ExchangeRateResponse exchangeRateResponse) {
        System.out.println(EXCHANGE_RATE_TITLE_MESSAGE);
        String currencyUnit = exchangeRateResponse.currencyUnit();
        String ttsRate = exchangeRateResponse.transferSelling();
        System.out.printf(EXCHANGE_RATE_FORMAT, currencyUnit, ttsRate);
        System.out.println();
    }

    public void printPaymentResult(PartitionResponse partitionResponse) {
        List<String> goodsNames = partitionResponse.goodsList().stream()
                        .map(Goods::getName)
                        .toList();

        String joinedGoodsNames = String.join(JOIN_DELIMITER, goodsNames);
        String afterPrice = partitionResponse.afterPrice();

        System.out.printf(PAYMENT_SUCCESS_FORMAT, joinedGoodsNames, afterPrice);
        System.out.println();
    }
}
