using System;

namespace ConsoleApp2
{
    class Program
    {
        static void Main(string[] args)
        {
            int num, numhoras;
			double valohora;
			
			Console.WriteLine("Informe o número: ");
			num = int.Parse(Console.ReadLine());
		    Console.WriteLine("Numero de horas trabalhadas: "); 			
            numhoras = int.Parse(Console.ReadLine());
			Console.WriteLine("valor hora: ");			
            valohora = int.Parse(Console.ReadLine());
			
			Console.WriteLine("Num: " + num + " Salary: "+ numhoras * valohora);	
			
        }
    }
}
