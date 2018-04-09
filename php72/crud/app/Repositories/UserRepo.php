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

    public function getAll(): array
    {
        return $this->model->get()->toArray();
    }

    public function store(array $data): array
    {
        return $this->model->create($data)->toArray();
    }
}