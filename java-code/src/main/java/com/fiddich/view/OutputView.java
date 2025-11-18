package com.fiddich.view;

import com.fiddich.model.DiscountInfoResponse;
import com.fiddich.service.DiscountPolicy;

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
}
