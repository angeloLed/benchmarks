using System.Threading.Tasks;
using Microsoft.AspNetCore.Mvc;
using App.Interfaces;
using App.Models;
using App.Infrastructure;
using System;
using System.Collections.Generic;

namespace App.Controllers
{
    [Produces("application/json")]
    [Route("")]
    public class HeatsController : Controller
    {
        private readonly IHeatRepository _heatRepository;

        public HeatsController(IHeatRepository heatRepository)
        {
            _heatRepository = heatRepository;
        }

        [NoCache]
        [HttpGet("heats")]
        public async Task<IEnumerable<Heat>> Get()
        {
            return await _heatRepository.GetAllHeats();
        }

        // GET userHeats
        [NoCache]
        [HttpGet("userHeats")]
        public async Task<IEnumerable<Heat>> getAllUserHasHeatZone(int x, int y, int radius)
        {
            return await _heatRepository.getAllUserHasHeatZone(x, y, radius);
        }

        // GET heats/:id
        [HttpGet("heats/{id}")]
        public async Task<Heat> Get(string id)
        {
            return await _heatRepository.GetHeat(id) ?? new Heat();
        }

        // POST heats
        [HttpPost("heats")]
        public void Post([FromBody] HeatParam newHeat)
        {
            _heatRepository.AddHeat(new Heat
            {
                Id = newHeat.Id,
                X = newHeat.X,
                Y = newHeat.Y,
                User = newHeat.User
            });
        }

        // PUT heats/:id
        // [HttpPut("{id}")]
        // public void Put(string id, [FromBody]string value)
        // {
        //     _heatRepository.UpdateHeatDocument(id, value);
        // }

        // // DELETE heats/:id
        // [HttpDelete("{id}")]
        // public void Delete(string id)
        // {
        //     _heatRepository.RemoveHeat(id);
        // }
    }
}
