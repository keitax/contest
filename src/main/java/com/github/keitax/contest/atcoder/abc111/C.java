package com.github.keitax.contest.atcoder.abc111;

import java.util.*;
import java.util.stream.Collectors;

public class C {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        int n = Integer.parseInt(sc.next());
        List<Integer> vs = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            vs.add(Integer.parseInt(sc.next()));
        }

        List<Integer> evens = new ArrayList<>();
        List<Integer> odds = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            if (i % 2 == 0) {
                evens.add(vs.get(i));
            } else {
                odds.add(vs.get(i));
            }
        }

        List<Integer> me = mode(evens);
        List<Integer> mo = mode(odds);

        int ans;
        if (me.get(0) != mo.get(0)) {
            int cnt = 0;
            for (int e : evens) {
                if (e != me.get(0)) {
                    cnt++;
                }
            }
            for (int o : odds) {
                if (o != mo.get(0)) {
                    cnt++;
                }
            }
            ans = cnt;
        } else {
            int cnt0 = 0;
            for (int e : evens) {
                if (e != me.get(0)) {
                    cnt0++;
                }
            }
            for (int o : odds) {
                if (mo.size() < 2 || o != mo.get(1)) {
                    cnt0++;
                }
            }
            int cnt1 = 0;
            for (int e : evens) {
                if (me.size() < 2 || e != me.get(1)) {
                    cnt1++;
                }
            }
            for (int o : odds) {
                if (o != mo.get(0)) {
                    cnt1++;
                }
            }
            ans = Math.min(cnt0, cnt1);
        }

        System.out.println(ans);
    }

    private static List<Integer> mode(List<Integer> xs) {
        Map<Integer, Integer> t = new HashMap<>();
        for (int x : xs) {
            Integer cnt = t.get(x);
            if (cnt == null) {
                cnt = 0;
            }
            t.put(x, cnt + 1);
        }
        List<Map.Entry<Integer, Integer>> es = new ArrayList(t.entrySet());
        es.sort((l, r) -> r.getValue() - l.getValue());
        return es.stream().map(Map.Entry::getKey).collect(Collectors.toList());

    }
}
