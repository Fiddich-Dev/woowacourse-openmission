package com.fiddich;

import com.fiddich.controller.PurchaseController;
import com.fiddich.view.InputParser;
import com.fiddich.view.InputView;
import com.fiddich.view.OutputView;

public class Application {
    public static void main(String[] args) {
        InputView inputView = new InputView();
        InputParser inputParser = new InputParser();
        OutputView outputView = new OutputView();
        PurchaseController purchaseController = new PurchaseController(inputView, inputParser, outputView);

        purchaseController.run();
    }
}