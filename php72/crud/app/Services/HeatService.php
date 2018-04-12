<?php

namespace App\Services;

use Illuminate\Http\Response;
use App\Repositories\HeatRepo;

/**
 * Class ApiController
 *
 * @package App\Http\Controllers
 * @version 1.0
 * @author Angelo (LED) <angelolandino@htomail.it>
 * @copyright Copyright (c) 2015-present
 */
class HeatService
{
    function __construct(HeatRepo $heatRepo)
    {
        $this->repo = $heatRepo;
    }

    public function getAll($filters = []): array
    {
        return $this->repo->getAll($filters);
    }

    public function getAllUserHasHeatZone($parameters): array
    {
         $filters = [];

        if( isset($parameters['x'])) {
            $minx = $parameters['x'] - ( $parameters['radius'] ?? 0);
            $maxx = $parameters['x'] + ( $parameters['radius'] ?? 0);
            $filters['x'] = [ '$gte' => $minx, '$lte' => $maxx ];
        }
        if( isset($parameters['y'])) {
            $miny = $parameters['y'] - ( $parameters['radius'] ?? 0);
            $maxy = $parameters['y'] + ( $parameters['radius'] ?? 0);
            $filters['y'] = [ '$gte' => $miny, '$lte' => $maxy ];
        }

        $heats = $this->repo->whereRaw($filters);
        $users = [];
        array_map( function($item) use (&$users) {
            if(! in_array($item['user'], $users)) {
                $users[] = $item['user'];
            }
        }, $heats);

        return $users;
    }

    public function store(array $data): array
    {
        return $this->repo->store($data);
    }
}