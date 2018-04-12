using Microsoft.Extensions.Options;
using MongoDB.Driver;
using App.Models;

namespace App.Data
{
    public class HeatContext
    {
        private readonly IMongoDatabase _database = null;

        public HeatContext(IOptions<Settings> settings)
        {
            var client = new MongoClient(settings.Value.ConnectionString);
            if (client != null)
                _database = client.GetDatabase(settings.Value.Database);
        }

        public IMongoCollection<Heat> Heats
        {
            get
            {
                return _database.GetCollection<Heat>("usersheats");
            }
        }
    }
}