
using MongoDB.Bson;

namespace App.Models
{
    public class HeatParam
    {
        public ObjectId Id { get; set; }
        public int X { get; set; } = 0;
        public int Y { get; set; } = 0;
        public string User { get; set; } = string.Empty;
    }
}