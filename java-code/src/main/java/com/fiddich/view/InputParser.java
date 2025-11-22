package com.fiddich.view;

import com.fiddich.Config;
import com.fiddich.model.Goods;

import java.math.BigDecimal;
import java.util.Arrays;

public class InputParser {

    private static final String INVALID_LONG_MESSAGE = "정수 입력이 올바르지 않습니다.";
    private static final String INVALID_GOODS_MESSAGE = "입력이 올바르지 않습니다.";

    public Long stringToLong(String input) {
        try {
            return Long.parseLong(input);
        } catch (NumberFormatException e) {
            throw new IllegalArgumentException(INVALID_LONG_MESSAGE);
        }
    }

    public Goods toGoods(String input) {
        String[] parts = input.trim().split("\\s+");
        if (parts.length != 2) {
            throw new IllegalArgumentException(INVALID_GOODS_MESSAGE);
        }
        String name = parts[0];
        String price = parts[1];

        return new Goods(name, price);
    }

    public boolean isConfirmed(String input) {
        if(input.equals(Config.PAYMENT_CONFIRM_MESSAGE)) {
            return true;
        }
        return false;
    }
}
