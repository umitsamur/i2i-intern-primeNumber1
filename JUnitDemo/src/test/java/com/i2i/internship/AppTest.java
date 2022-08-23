package com.i2i.internship;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.Test;


public class AppTest 
{
	@Test
	public void testPrimePositiveNumber() {
		boolean actual = App.isPrime(11);
		boolean expected = true;
		assertEquals(expected, actual);
	}
	
	@Test
	public void testNotPrimePositiveNumber() {
		boolean actual = App.isPrime(20);
		boolean expected = false;
		assertEquals(expected, actual);
	}
	
	@Test
	public void testIsPrimeNegativeNumber() {
		boolean actual = App.isPrime(-50);
		boolean expected = false;
		assertEquals(expected, actual);
	}
}
