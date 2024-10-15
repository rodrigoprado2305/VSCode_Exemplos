using System;

namespace ConsoleApp2
{
    class Program
    {
        static void Main(string[] args)
        {
            double pi = 3.14159;
            double raio, area;

            raio = double.Parse(Console.ReadLine());
            
            area = pi * Math.Pow(raio, 2);

            Console.WriteLine("Area = " + area);
        }
    }
}
