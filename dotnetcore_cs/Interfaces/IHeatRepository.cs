using System.Collections.Generic;
using System.Threading.Tasks;
using App.Models;

namespace App.Interfaces
{
    public interface IHeatRepository
    {
        Task<IEnumerable<Heat>> GetAllHeats();
        Task<IEnumerable<Heat>> getAllUserHasHeatZone(int x, int y, int radius);
        Task<Heat> GetHeat(string id);

        Task AddHeat(Heat item);

        // Task<bool> RemoveHeat(string id);

        // Task<bool> UpdateHeat(string id, string body);

        // Task<bool> UpdateHeatDocument(string id, string body);

        // Task<bool> RemoveAllHeats();

        Task<string> CreateIndex();
    }
}