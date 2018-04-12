<?php

namespace App\Transformers;

/**
 * Class Transformer
 *
 */
class HeatTransformer
{
    /**
     * Transform a collection to a response
     *
     * @param array $items
     * @return array
     */
    public function transformMany(array $items)
    {
        return array_map([$this, 'transform'], $items);
    }

    /**
     * To transform a single item.
     *
     * @param $item
     * @return mixed
     */
    public function transform($item): array
    {
        return [
            'heatId' => $item['_id'],
            'username' => $item['user'],
            'x' => $item['x'],
            'y' => $item['y']
        ];
    }
}
