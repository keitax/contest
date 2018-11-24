package com.github.keitax.contest.atcoder.abc113;

import java.util.Scanner;

public class A {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int x = Integer.parseInt(sc.next());
        int y = Integer.parseInt(sc.next());

        int ans = x + y / 2;

        System.out.println(ans);
    }
}
