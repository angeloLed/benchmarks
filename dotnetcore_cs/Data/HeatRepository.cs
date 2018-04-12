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

        public async Task<IEnumerable<Heat>> getAllUserHasHeatZone(int x, int y, int radius)
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

                return await _context.Heats.Find(BsonDocument.Parse("{"+query+"}")).ToListAsync();
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

        // public async Task<bool> RemoveHeat(string id)
        // {
        //     try
        //     {
        //         DeleteResult actionResult = await _context.Heats.DeleteOneAsync(
        //              Builders<Heat>.Filter.Eq("Id", id));

        //         return actionResult.IsAcknowledged 
        //             && actionResult.DeletedCount > 0;
        //     }
        //     catch (Exception ex)
        //     {
        //         // log or manage the exception
        //         throw ex;
        //     }
        // }

        // public async Task<bool> UpdateHeat(string id, string body)
        // {
        //     var filter = Builders<Heat>.Filter.Eq(s => s.Id, id);
        //     var update = Builders<Heat>.Update
        //                     .Set(s => s.Body, body)
        //                     .CurrentDate(s => s.UpdatedOn);

        //     try
        //     {
        //         UpdateResult actionResult = await _context.Heats.UpdateOneAsync(filter, update);

        //         return actionResult.IsAcknowledged
        //             && actionResult.ModifiedCount > 0;
        //     }
        //     catch (Exception ex)
        //     {
        //         // log or manage the exception
        //         throw ex;
        //     }
        // }

        // public async Task<bool> UpdateHeat(string id, Heat item)
        // {
        //     try
        //     {
        //         ReplaceOneResult actionResult = await _context.Heats
        //                                         .ReplaceOneAsync(n => n.Id.Equals(id)
        //                                                         , item
        //                                                         , new UpdateOptions { IsUpsert = true });
        //         return actionResult.IsAcknowledged
        //             && actionResult.ModifiedCount > 0;
        //     }
        //     catch (Exception ex)
        //     {
        //         // log or manage the exception
        //         throw ex;
        //     }
        // }

        // Demo function - full document update
        // public async Task<bool> UpdateHeatDocument(string id, string body)
        // {
        //     var item = await GetHeat(id) ?? new Heat();
        //     item.Body = body;
        //     item.UpdatedOn = DateTime.Now;

        //     return await UpdateHeat(id, item);
        // }

        // public async Task<bool> RemoveAllHeats()
        // {
        //     try
        //     {
        //         DeleteResult actionResult = await _context.Heats.DeleteManyAsync(new BsonDocument());

        //         return actionResult.IsAcknowledged
        //             && actionResult.DeletedCount > 0;
        //     }
        //     catch (Exception ex)
        //     {
        //         // log or manage the exception
        //         throw ex;
        //     }
        // }

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