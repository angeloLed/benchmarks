using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using Microsoft.Extensions.Options;
using MongoDB.Driver;

using App.Interfaces;
using App.Models;
using MongoDB.Bson;

namespace App.Data
{
    public class HeatRepository : IHeatRepository
    {
        private readonly HeatContext _context = null;

        public HeatRepository(IOptions<Settings> settings)
        {
            _context = new HeatContext(settings);
        }

        public async Task<IEnumerable<Heat>> GetAllHeats()
        {
            try
            {
                return await _context.Heats.Find(_ => true).ToListAsync();
            }
            catch (Exception ex)
            {
                // log or manage the exception
                throw ex;
            }
        }

        public async Task<List<string>> getAllUserHasHeatZone(int x, int y, int radius)
        {
            try
            {
                string query = "";

                int min = x - radius;
                int max = x + radius;
                query += " x: { $gte: "+min+", $lte: "+max+" },";

                min = y - radius;
                max = y + radius;

                query += " y: { $gte: "+min+", $lte: "+max+" }";

                IEnumerable<Heat> heats = await _context.Heats.Find(BsonDocument.Parse("{"+query+"}")).ToListAsync();

                List<string> users = new List<string>();

                foreach(Heat heat in heats)
                {
                    if(! users.Contains(heat.User)) {
                        users.Add(heat.User);
                    }
                }

                return users;

            }
            catch (Exception ex)
            {
                // log or manage the exception
                throw ex;
            }
        }

        public async Task<Heat> GetHeat(string id)
        {
            try
            {
                ObjectId objId = GetInternalId(id);
                return await _context.Heats
                                .Find(Heat => Heat.Id == objId )
                                .FirstOrDefaultAsync();
            }
            catch (Exception ex)
            {
                // log or manage the exception
                throw ex;
            }
        }

        private ObjectId GetInternalId(string id)
        {
            ObjectId internalId;
            if (!ObjectId.TryParse(id, out internalId))
                internalId = ObjectId.Empty;

            return internalId;
        }

        public async Task AddHeat(Heat item)
        {
            try
            {
                await _context.Heats.InsertOneAsync(item);
            }
            catch (Exception ex)
            {
                // log or manage the exception
                throw ex;
            }
        }

        public async Task<string> CreateIndex()
        {
            try
            {
                return await _context.Heats.Indexes
                    .CreateOneAsync(Builders<Heat>
                        .IndexKeys
                        .Ascending(item => item.User)
                        .Ascending(item => item.X)
                        .Ascending(item => item.Y)
                );
            }
            catch (Exception ex)
            {
                // log or manage the exception
                throw ex;
            }
        }
    }
}