package com.fiddich.model;

public class Goods {
    private String name;
    private String price;

    public Goods() {
    }

    public Goods(String name, String price) {
        this.name = name;
        this.price = price;
    }

    public String getName() {
        return name;
    }

    public String getPrice() {
        return price;
    }
}
