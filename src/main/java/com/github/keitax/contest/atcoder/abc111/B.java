package com.github.keitax.contest.atcoder.abc111;

import java.util.Scanner;

public class B {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = Integer.parseInt(scanner.next());
        for (int i = n; ; i++) {
            if (isTargetContest(i)) {
                System.out.println(i);
                return;
            }
        }
    }

    private static boolean isTargetContest(int n) {
        int k0 = n % 10;
        for (n /= 10; n > 0; n /= 10) {
            if (k0 != n % 10) {
                return false;
            }
        }
        return true;
    }
}
