<?php

namespace App\Repositories;

use Illuminate\Http\Response;
use App\Models\User;

/**
 * Class ApiController
 *
 * @package App\Http\Controllers
 * @version 1.0
 * @author Angelo (LED) <angelolandino@htomail.it>
 * @copyright Copyright (c) 2015-present
 */
class UserRepo
{
    function __construct(User $user)
    {
        $this->model = $user;
    }

    public function getAll($filters = []): array
    {
        return $this->model->get($filters)->toArray();
    }

    public function whereRaw($filters = []): array
    {
        return $this->model->whereRaw($filters)->get()->toArray();
    }

    public function store(array $data): array
    {
        return $this->model->create($data)->toArray();
    }
}