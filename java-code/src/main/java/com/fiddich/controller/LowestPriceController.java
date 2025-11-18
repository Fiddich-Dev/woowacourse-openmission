package com.fiddich.controller;

import com.fiddich.model.DiscountInfoResponse;
import com.fiddich.model.Goods;
import com.fiddich.model.PartitionResponse;
import com.fiddich.model.ResponseFormat;
import com.fiddich.service.DiscountPolicy;
import com.fiddich.service.Partition;
import com.fiddich.view.InputParser;
import com.fiddich.view.InputView;
import com.fiddich.view.OutputView;

import java.util.ArrayList;
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
        DiscountPolicy discountPolicy = new DiscountPolicy();

        List<DiscountInfoResponse> couponDiscountInfoResponseList = discountPolicy.byCoupon().getContent();
        outputView.printCouponDiscountInfo(couponDiscountInfoResponseList);
        List<DiscountInfoResponse> cardDiscountInfoResponseList = discountPolicy.byCard().getContent();
        outputView.printCouponDiscountInfo(cardDiscountInfoResponseList);

        List<Goods> goodsList = inputGoodsList();

        Partition partition = new Partition();
        List<PartitionResponse> partitionResponseList = partition.apply(goodsList).getContent();

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