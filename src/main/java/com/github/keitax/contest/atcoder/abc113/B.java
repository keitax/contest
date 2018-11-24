package com.github.keitax.contest.atcoder.abc113;

import java.util.Scanner;

public class B {
    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int N = Integer.parseInt(scanner.next());
        int T = Integer.parseInt(scanner.next());
        int A = Integer.parseInt(scanner.next());
        int[] H = new int[N];
        for (int i = 0; i < N; i++) {
            H[i] = Integer.parseInt(scanner.next());
        }

        double min = Integer.MAX_VALUE;
        int ans = 1;
        for (int i = 0; i < N; i++) {
            double d = Math.abs(A - (T - H[i] * 0.006));
            if (d < min) {
                min = d;
                ans = i + 1;
            }
        }

        System.out.println(ans);
    }
}
