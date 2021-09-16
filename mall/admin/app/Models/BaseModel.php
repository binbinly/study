<?php


namespace App\Models;


use Illuminate\Database\Eloquent\Model;
use Illuminate\Support\Facades\DB;

/**
 * 基础模型
 * Class BaseModel
 * @package App\Models
 */
class BaseModel extends Model
{
    const CONN_DEFAULT = '';
    const CONN_MALL = 'connection_mall';

    protected $conn = self::CONN_MALL;

    protected $table = '';

    public function __construct(array $attributes = [])
    {
        $this->setConnection(self::connection($this->conn));

        $this->setTable($this->table);

        parent::__construct($attributes);
    }

    protected $dateFormat = '';

    protected $dates = [
        'created_at',
        'updated_at',
        'deleted_at',
        'start_at',
        'end_at'
    ];

    /**
     * 时间戳存储
     * @param  $value
     * @return false|int|string
     */
    public function fromDateTime($value)
    {
        return strtotime(parent::fromDateTime($value));
    }

    /**
     * 事务开启
     * @param string $conn
     */
    public static function transaction($conn = self::CONN_MALL)
    {
        DB::connection(self::connection($conn))->beginTransaction();
    }

    /**
     * 事务提交
     * @param string $conn
     */
    public static function commit($conn = self::CONN_MALL)
    {
        DB::connection(self::connection($conn))->commit();
    }

    /**
     * 事务回滚
     * @param string $conn
     */
    public static function rollBack($conn = self::CONN_MALL)
    {
        DB::connection(self::connection($conn))->rollBack();
    }

    public static function selectMy($sql, $conn = self::CONN_MALL)
    {
        return DB::connection(self::connection($conn))->select($sql);
    }

    /**
     * 有些数据库语句不会返回任何值。对于这些语句，可以在 DB facade 上使用 statement 方法来操作：
     * @param $sql
     * @param string $conn
     */
    public static function statementMy($sql, $conn = self::CONN_MALL)
    {
        DB::connection(self::connection($conn))->statement($sql);
    }

    /**
     * 获取连接
     * @param $conn
     * @return mixed
     */
    protected static function connection($conn)
    {
        if ($conn) {
            $connection = config('admin.database.' . $conn);
        } else {
            $connection = config('database.default');
        }
        return $connection;
    }
}