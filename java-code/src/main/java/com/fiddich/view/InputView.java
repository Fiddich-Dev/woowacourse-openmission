package com.fiddich.view;

import com.fiddich.config.Config;
import com.fiddich.util.Console;

import java.util.ArrayList;
import java.util.List;

public class InputView {

    private static final String REQUEST_GOODS_COUNT_MESSAGE = "구매하고 싶은 상품의 개수를 입력해주세요.";
    private static final String REQUEST_GOODS_MESSAGE = "구매하고 싶은 상품의 이름과 가격(달러)을 입력해주세요.";
    private static final String REQUEST_CONFIRM_MESSAGE = "결제 하시러면 %s를 입력해주세요%n";

    public String requestGoodsCount() {
        System.out.println(REQUEST_GOODS_COUNT_MESSAGE);
        String input = Console.readLine();
        System.out.println();
        return input;
    }

    public List<String> requestGoods(long count) {
        System.out.println(REQUEST_GOODS_MESSAGE);
        List<String> result = new ArrayList<>();
        for(long i = 0; i < count; i++) {
            result.add(Console.readLine());
        }
        System.out.println();
        return result;
    }

    public String requestConfirmPayment() {
        System.out.printf(REQUEST_CONFIRM_MESSAGE, Config.PAYMENT_CONFIRM_MESSAGE);
        String input = Console.readLine();
        System.out.println();
        return input;
    }
}
