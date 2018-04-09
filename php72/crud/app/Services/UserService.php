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

    public function getAll(): array
    {
        return $this->repo->getAll();
    }

    public function store(array $data): array
    {
        return $this->repo->store($data);
    }
}