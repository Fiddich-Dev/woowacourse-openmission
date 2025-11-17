package com.fiddich.model;

import java.math.BigDecimal;

public class Goods {
    private final String name;
    private final BigDecimal price;

    public Goods(String name, BigDecimal price) {
        this.name = name;
        this.price = price;
    }

    public String getName() {
        return name;
    }

    public BigDecimal getPrice() {
        return price;
    }
}
