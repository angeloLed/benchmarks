<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use Illuminate\Http\Response;
use Illuminate\Http\JsonResponse;

use App\Transformers\HeatTransformer;
use App\Services\HeatService;

class HeatController extends Controller
{
    /**
     * Create a new controller instance.
     *
     * @return void
     */
    public function __construct(HeatService $heatService, HeatTransformer $heatTransformer)
    {
        $this->service = $heatService;
        $this->transformer = $heatTransformer;
    }

    protected function respond($data, $headers = [])
    {
        return response()->json($data, $this->statusCode, $headers);
    }

    public function all(): JsonResponse
    {
        $heats = $this->service->getAll();
        $this->statusCode = 200;
        return $this->respond(
            ['data' => $this->transformer->transformMany($heats)]
        );
    }

    public function getAllUserHasHeatZone(Request $request): JsonResponse
    {
        $users = $this->service->getAllUserHasHeatZone($request->all());
        $this->statusCode = 200;
        return $this->respond(
            ['data' => $users]
        );
    }

    public function store(Request $request): JsonResponse
    {
        $heat = $this->service->store($request->all());
        $this->statusCode = 201;
        return $this->respond(
            ['data' => $this->transformer->transform($heat)]
        );
    }
}
