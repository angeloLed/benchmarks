using System.Collections.Generic;
using System.Threading.Tasks;
using App.Models;

namespace App.Interfaces
{
    public interface IHeatRepository
    {
        Task<IEnumerable<Heat>> GetAllHeats();
        Task<List<string>> getAllUserHasHeatZone(int x, int y, int radius);
        Task<Heat> GetHeat(string id);

        Task AddHeat(Heat item);

        Task<string> CreateIndex();
    }
}