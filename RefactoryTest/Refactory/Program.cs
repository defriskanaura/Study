//using Newtonsoft.Json;
using System;
using System.Collections.Generic;

public class Program
{
	static void Main()	
	{
		string word = "";
		Console.WriteLine("nomor satu");
		word = Console.ReadLine();
		string reverse = "";
		for (int i = word.Length-1; i>=0; i--)
		{
			reverse += word[i].ToString();
		}
		if (reverse == word)
		{
			Console.WriteLine("true");
		}
		else
		{
			Console.WriteLine("false");
		}
		//Console.ReadKey();
		
		Console.WriteLine("nomor dua");
		string a = "abc";
		Console.WriteLine(a.ToLower());
		string b = "ABC";
		Console.WriteLine(b.ToUpper());
		string c = "h";
		string d = "ello";
		string e = "w";
		string f = "orld";
		Console.WriteLine("{0}{1} {2}{3}", c.ToLower(), d.ToUpper(), e.ToLower(), f.ToUpper());
	
		Console.WriteLine("nomor tiga");
		Console.WriteLine("kata satu");
		string word1 = "";
		word1 = Console.ReadLine();
		Console.WriteLine("kata dua");
		string word2 = fi"";
		word2 = Console.ReadLine();
		char[] char1 = word1.ToLower().ToCharArray();
		char[] char2 = word2.ToLower().ToCharArray();
		
		Array.Sort(char1);
		Array.Sort(char2);
		
		string word3 = new string(char1);
		string word4 = new string(char2);
		if(word3 == word4)
		{
			Console.WriteLine("true");
		}
		else
		{
			Console.WriteLine("false");
		}
	}
}