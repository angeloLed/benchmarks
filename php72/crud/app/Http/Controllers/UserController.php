<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Illuminate\Http\JsonResponse;

use App\Transformers\UserTransformer;
use App\Services\UserService;

class UserController extends Controller
{
    /**
     * Create a new controller instance.
     *
     * @return void
     */
    public function __construct(UserService $userService, UserTransformer $userTransformer)
    {
        $this->service = $userService;
        $this->transformer = $userTransformer;
    }

    protected function respond($data, $headers = [])
    {
        return response()->json($data, $this->statusCode, $headers);
    }

    public function all(): JsonResponse
    {
        $users = $this->service->getAll();
        $this->statusCode = 200;
        return $this->respond(
            ['data' => $this->transformer->transformMany($users)]
        );
    }

    public function store(Request $request): JsonResponse
    {
        $user = $this->service->store($request->all());
        $this->statusCode = 201;
        return $this->respond(
            ['data' => $this->transformer->transform($user)]
        );
    }
}
