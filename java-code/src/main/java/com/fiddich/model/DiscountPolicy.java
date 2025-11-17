package com.fiddich.model;

import java.math.BigDecimal;

public interface DiscountPolicy {
    BigDecimal apply(BigDecimal price);
}
