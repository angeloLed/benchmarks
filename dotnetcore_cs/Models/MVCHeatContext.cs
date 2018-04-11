    using Microsoft.EntityFrameworkCore;  
      
    namespace App.Models  
    {  
        public class MvcHeatContext : DbContext  
        {  
            public MvcHeatContext (DbContextOptions<MvcHeatContext> options)  
                : base(options)  
            {  
            }  
      
            public DbSet<App.Models.Heat> Heat { get; set; }  
        }  
    }  
