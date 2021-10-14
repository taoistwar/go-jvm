package com.github.taoistwar.java;

import java.io.Serializable;

public class GaussTest implements Serializable {
  
    public static void main(String[] args) throws RuntimeException {
        int sum = 0;
        for (int i=0; i<100; i++) {
            sum += i;
        }
        System.out.println(sum);
    }


}