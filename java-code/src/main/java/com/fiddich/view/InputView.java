package com.fiddich.view;

import com.fiddich.Console;

import java.util.ArrayList;
import java.util.List;

public class InputView {

    public String requestAmount() {
        System.out.println("현재 잔액을 입력하세요.(원화)");
        String input = Console.readLine();
        System.out.println();
        return input;
    }

    public String requestGoodsCount() {
        System.out.println("구매하고 싶은 상품의 개수를 입력해주세요.");
        String input = Console.readLine();
        System.out.println();
        return input;
    }

    public List<String> requestGoods(long count) {
        System.out.println("구매하고 싶은 상품의 이름과 가격(달러)을 소수점 둘째자리까지 입력해주세요.");
        List<String> result = new ArrayList<>();
        for(long i = 0; i < count; i++) {
            result.add(Console.readLine());
        }
        System.out.println();
        return result;
    }


}
