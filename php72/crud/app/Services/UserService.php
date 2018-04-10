<?php

namespace App\Services;

use Illuminate\Http\Response;
use App\Repositories\UserRepo;

/**
 * Class ApiController
 *
 * @package App\Http\Controllers
 * @version 1.0
 * @author Angelo (LED) <angelolandino@htomail.it>
 * @copyright Copyright (c) 2015-present
 */
class UserService
{
    function __construct(UserRepo $userRepo)
    {
        $this->repo = $userRepo;
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
        if( isset($parameters['user'])) {
            $filters['user'] = $parameters['user'];
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