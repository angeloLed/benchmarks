using System;
using MongoDB.Bson;
using MongoDB.Bson.Serialization.Attributes;
    
namespace App.Models
{  
    public class Heat  
    {  
        [BsonId]
        public ObjectId Id { get; set; }
        [BsonElement("user")]
        public string User { get; set; }
        [BsonElement("x")]
        public int X { get; set; }
        [BsonElement("y")]
        public int Y { get; set; }
    }  
}  
