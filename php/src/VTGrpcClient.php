<?php
require_once (dirname(__FILE__) . '/proto/vtgateservice.php');

require_once (dirname(__FILE__) . '/VTRpcClient.php');

class VTGrpcClient implements VTRpcClient {
	protected $stub;

	public function __construct($addr, $opts = []) {
		$this->stub = new Proto\Vtgateservice\VitessClient($addr, $opts);
	}

	public function execute(VTContext $ctx, Proto\Vtgate\ExecuteRequest $request) {
		list ( $response, $status ) = $this->stub->Execute($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function executeShards(VTContext $ctx, Proto\Vtgate\ExecuteShardsRequest $request) {
		list ( $response, $status ) = $this->stub->ExecuteShards($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function executeKeyspaceIds(VTContext $ctx, Proto\Vtgate\ExecuteKeyspaceIdsRequest $request) {
		list ( $response, $status ) = $this->stub->ExecuteKeyspaceIds($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function executeKeyRanges(VTContext $ctx, Proto\Vtgate\ExecuteKeyRangesRequest $request) {
		list ( $response, $status ) = $this->stub->ExecuteKeyRanges($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function executeEntityIds(VTContext $ctx, Proto\Vtgate\ExecuteEntityIdsRequest $request) {
		list ( $response, $status ) = $this->stub->ExecuteEntityIds($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function executeBatchShards(VTContext $ctx, Proto\Vtgate\ExecuteBatchShardsRequest $request) {
		list ( $response, $status ) = $this->stub->ExecuteBatchShards($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function executeBatchKeyspaceIds(VTContext $ctx, Proto\Vtgate\ExecuteBatchKeyspaceIdsRequest $request) {
		list ( $response, $status ) = $this->stub->ExecuteBatchKeyspaceIds($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function streamExecute(VTContext $ctx, Proto\Vtgate\StreamExecuteRequest $request) {
		return new VTGrpcStreamResponse($this->stub->StreamExecute($request));
	}

	public function streamExecuteShards(VTContext $ctx, Proto\Vtgate\StreamExecuteShardsRequest $request) {
		return new VTGrpcStreamResponse($this->stub->StreamExecuteShards($request));
	}

	public function streamExecuteKeyspaceIds(VTContext $ctx, Proto\Vtgate\StreamExecuteKeyspaceIdsRequest $request) {
		return new VTGrpcStreamResponse($this->stub->StreamExecuteKeyspaceIds($request));
	}

	public function streamExecuteKeyRanges(VTContext $ctx, Proto\Vtgate\StreamExecuteKeyRangesRequest $request) {
		return new VTGrpcStreamResponse($this->stub->StreamExecuteKeyRanges($request));
	}

	public function begin(VTContext $ctx, Proto\Vtgate\BeginRequest $request) {
		list ( $response, $status ) = $this->stub->Begin($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function commit(VTContext $ctx, Proto\Vtgate\CommitRequest $request) {
		list ( $response, $status ) = $this->stub->Commit($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function rollback(VTContext $ctx, Proto\Vtgate\RollbackRequest $request) {
		list ( $response, $status ) = $this->stub->Rollback($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function getSrvKeyspace(VTContext $ctx, Proto\Vtgate\GetSrvKeyspaceRequest $request) {
		list ( $response, $status ) = $this->stub->GetSrvKeyspace($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function splitQuery(VTContext $ctx, Proto\Vtgate\SplitQueryRequest $request) {
		list ( $response, $status ) = $this->stub->SplitQuery($request)->wait();
		self::checkError($status);
		return $response;
	}

	public function close() {
		$this->stub->close();
	}

	public static function checkError($status) {
		if ($status) {
			switch ($status->code) {
				case Grpc\STATUS_OK:
					break;
				case Grpc\STATUS_INVALID_ARGUMENT:
					throw new VTBadInputError($status->details);
				case Grpc\STATUS_DEADLINE_EXCEEDED:
					throw new VTDeadlineExceededError($status->details);
				case Grpc\STATUS_ALREADY_EXISTS:
					throw new VTIntegrityError($status->details);
				case Grpc\STATUS_UNAUTHENTICATED:
					throw new VTUnauthenticatedError($status->details);
				case Grpc\STATUS_UNAVAILABLE:
					throw new VTTransientError($status->details);
				default:
					throw new VTException("{$status->code}: {$status->details}");
			}
		}
	}
}

class VTGrpcStreamResponse {
	private $call;
	private $iterator;

	public function __construct($call) {
		$this->call = $call;
		$this->iterator = $call->responses();
		
		if (! $this->iterator->valid()) {
			// No responses were returned. Check for error.
			VTGrpcClient::checkError($this->call->getStatus());
		}
	}

	public function next() {
		if ($this->iterator->valid()) {
			$value = $this->iterator->current();
			$this->iterator->next();
			return $value;
		} else {
			// Now that the responses are done, check for gRPC status.
			VTGrpcClient::checkError($this->call->getStatus());
			// If no exception is raised, indicate successful end of stream.
			return FALSE;
		}
	}

	public function close() {
		$this->call->cancel();
	}
}
