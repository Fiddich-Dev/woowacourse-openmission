package com.fiddich.view;

import com.fiddich.model.Goods;

import java.math.BigDecimal;
import java.util.Arrays;

public class InputParser {

    public Long stringToLong(String input) {
        try {
            return Long.parseLong(input);
        } catch (NumberFormatException e) {
            throw new IllegalArgumentException("정수 입력이 올바르지 않습니다.");
        }
    }

    public BigDecimal StringToBigDecimal(String input) {
        try {
            return new BigDecimal(input);
        } catch (RuntimeException e) {
            throw new IllegalArgumentException("유리수 입력이 올바르지 않습니다.");
        }
    }

    public Goods toGoods(String input) {
        String[] parts = input.trim().split("\\s+");
        if (parts.length != 2) {
            throw new IllegalArgumentException("입력이 올바르지 않습니다.");
        }
        String name = parts[0];
        BigDecimal price = StringToBigDecimal(parts[1]);

        return new Goods(name, price);
    }
}
