package com.fiddich.model;

import java.math.BigDecimal;
import java.util.Arrays;
import java.util.Comparator;
import java.util.List;

public enum CardDiscountPolicyByDollar implements DiscountPolicy {
    OVER_50_MINUS_8(
            new BigDecimal(50),
            new BigDecimal(8)
    ),
    OVER_100_MINUS_16(
            new BigDecimal(100),
            new BigDecimal(16)
    ),
    OVER_150_MINUS_30(
            new BigDecimal(150),
            new BigDecimal(30)
    ),
    OVER_300_MINUS_60(
            new BigDecimal(300),
            new BigDecimal(60)
    );

    public static final Comparator<CardDiscountPolicyByDollar> BY_DISCOUNT_AMOUNT_ASC = (p1, p2) ->
            p1.discountAmount.compareTo(p2.discountAmount);
    public static final Comparator<CardDiscountPolicyByDollar> BY_DISCOUNT_AMOUNT_DESC = (p1, p2) ->
            p2.discountAmount.compareTo(p1.discountAmount);

    private final BigDecimal threshold;
    private final BigDecimal discountAmount;

    CardDiscountPolicyByDollar(BigDecimal threshold, BigDecimal discountAmount) {
        this.threshold = threshold;
        this.discountAmount = discountAmount;
    }

    public BigDecimal getDiscountAmount() {
        return discountAmount;
    }

    public BigDecimal getThreshold() {
        return threshold;
    }

    public List<CardDiscountPolicyByDollar> getValues() {
        return Arrays.stream(values())
                        .sorted(BY_DISCOUNT_AMOUNT_ASC)
                        .toList();
    }

    public List<CardDiscountPolicyByDollar> getValues(Comparator<CardDiscountPolicyByDollar> comparator) {
        return Arrays.stream(values())
                .sorted(comparator)
                .toList();
    }

    @Override
    public BigDecimal apply(BigDecimal price) {
        return getValues(BY_DISCOUNT_AMOUNT_DESC)
                .stream()
                .filter(policy -> price.compareTo(policy.threshold) >= 0)
                .findFirst()
                .map(policy -> price.subtract(policy.discountAmount))
                .orElse(price);
    }
}
