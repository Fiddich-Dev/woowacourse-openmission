package com.fiddich.view;

import com.fiddich.model.*;
import com.fiddich.model.dto.resonse.DiscountInfoResponse;
import com.fiddich.model.dto.resonse.ExchangeRateResponse;
import com.fiddich.model.dto.resonse.PartitionResponse;

import java.util.List;

public class OutputView {

    public void printCouponDiscountInfo(List<DiscountInfoResponse> discountInfoResponseList) {
        System.out.println("쿠폰 할인");
        printDiscountInfo(discountInfoResponseList);
        System.out.println();
    }

    public void printCardDiscountInfo(List<DiscountInfoResponse> discountInfoResponseList) {
        System.out.println("카드 할인");
        printDiscountInfo(discountInfoResponseList);
        System.out.println();
    }

    private void printDiscountInfo(List<DiscountInfoResponse> discountInfoResponseList) {
        for (DiscountInfoResponse discountInfoResponse : discountInfoResponseList) {
            System.out.println(discountInfoResponse.threshold() + " 이상 구매하면 " + discountInfoResponse.discountAmount() + " 할인");
        }
    }

    public void printDiscount(List<PartitionResponse> partitionResponseList) {
        for(PartitionResponse partitionResponse : partitionResponseList) {
            List<String> goodsNames = partitionResponse.goodsList().stream()
                    .map(Goods::getName)
                    .toList();

            System.out.println(String.join(", ", goodsNames) + " : " + partitionResponse.beforePrice() + " -> " + partitionResponse.afterPrice());

        }
    }

    public void printExchangeRate(ExchangeRateResponse exchangeRateResponse) {
        System.out.println("현재 환율");
        System.out.println(1 + exchangeRateResponse.currencyUnit() + "는 " + exchangeRateResponse.transferSelling() + "원");
        System.out.println();
    }
}
