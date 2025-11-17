package com.fiddich.model;

import java.math.BigDecimal;
import java.util.Arrays;
import java.util.Comparator;
import java.util.List;

public enum CouponDiscountPolicyByDollar implements DiscountPolicy {
    OVER_25_MINUS_4(
            new BigDecimal(25),
            new BigDecimal(4)
    ),
    OVER_54_MINUS_8(
            new BigDecimal(54),
            new BigDecimal(8)
    ),
    OVER_89_MINUS_13(
            new BigDecimal(89),
            new BigDecimal(13)
    ),
    OVER_209_MINUS_31(
            new BigDecimal(209),
            new BigDecimal(31)
    ),
    OVER_679_MINUS_102(
            new BigDecimal(679),
            new BigDecimal(102)
    );

    public static final Comparator<CouponDiscountPolicyByDollar> BY_DISCOUNT_AMOUNT_ASC = (p1, p2) ->
            p1.discountAmount.compareTo(p2.discountAmount);
    public static final Comparator<CouponDiscountPolicyByDollar> BY_DISCOUNT_AMOUNT_DESC = (p1, p2) ->
            p2.discountAmount.compareTo(p1.discountAmount);

    private final BigDecimal threshold;
    private final BigDecimal discountAmount;

    CouponDiscountPolicyByDollar(BigDecimal threshold, BigDecimal discountAmount) {
        this.threshold = threshold;
        this.discountAmount = discountAmount;
    }

    public BigDecimal getThreshold() {
        return threshold;
    }

    public BigDecimal getDiscountAmount() {
        return discountAmount;
    }

    public List<CouponDiscountPolicyByDollar> getValues() {
        return Arrays.stream(values())
                .sorted(BY_DISCOUNT_AMOUNT_ASC)
                .toList();
    }

    public List<CouponDiscountPolicyByDollar> getValues(Comparator<CouponDiscountPolicyByDollar> comparator) {
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
