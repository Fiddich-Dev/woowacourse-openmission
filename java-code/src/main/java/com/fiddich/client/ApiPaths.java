package com.fiddich.client;

public class ApiPaths {
    public static final String BASE_URL = "http://localhost:8080";

    public static final String COUPON_DISCOUNT_POLICY = "/coupon-discount-policy";
    public static final String CARD_DISCOUNT_POLICY = "/card-discount-policy";
    public static final String EXCHANGE_RATE = "/exchange-rate";
    public static final String GOODS_PARTITION = "/partition";
    public static final String PAYMENT = "/payment";

    public static String createPath(String path) {
        return BASE_URL + path;
    }
}
