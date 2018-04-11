using System;  
using System.ComponentModel.DataAnnotations;  
    
namespace App.Models  
{  
    public class Heat  
    {  
        public int _Id { get; set; }  
        [Required]  
        public string User { get; set; }  
        [Required]  
        public int X { get; set; }  
        [Required]  
        public int Y { get; set; }  
    }  
}  
