package com.fiddich.controller;

import com.fiddich.model.Goods;
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
        List<Goods> goodsList = inputGoodsList();



    }

    private List<Goods> inputGoodsList() {
        Long goodsCount = inputParser.stringToLong(inputView.requestGoodsCount());
        return inputView.requestGoods(goodsCount).stream()
                .map(inputParser::toGoods)
                .toList();
    }
}
