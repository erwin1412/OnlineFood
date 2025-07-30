package transaction

import (
	"context"
	"transaction-service/internal/app"
	pb "transaction-service/internal/delivery/grpc/pb/transaction"
	"transaction-service/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type TransactionHandler struct {
	pb.UnimplementedTransactionServiceServer
	App *app.TransactionApp
}

func NewTransactionHandler(app *app.TransactionApp) *TransactionHandler {
	return &TransactionHandler{App: app}
}

func (h *TransactionHandler) CreateTransaction(ctx context.Context, req *pb.CreateTransactionRequest) (*pb.TransactionResponse, error) {
	// fmt.Println("CreateTransaction called with request:", req)

	tx := &domain.Transaction{
		UserID:     req.GetUserId(),
		CourierID:  req.GetCourierId(),
		MerchantID: req.GetMerchantId(),
		Status:     req.GetStatus(),
		SnapToken:  req.GetSnapToken(),
	}

	var details []*domain.TransactionDetail
	var total int64
	total = 0
	for _, d := range req.GetDetails() {
		details = append(details, &domain.TransactionDetail{
			FoodID: d.GetFoodId(),
			Qty:    d.GetQty(),
			Price:  d.GetPrice(),
		})
		total += d.GetQty() * d.GetPrice()
	}
	tx.Total = total

	createdTx, err := h.App.Create(ctx, tx, details, tx.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create transaction: %v", err)
	}

	return &pb.TransactionResponse{
		Transaction: mapToPbTransaction(createdTx),
	}, nil
}

// ✅ GetAllTransaction
func (h *TransactionHandler) GetAllTransaction(ctx context.Context, req *pb.GetAllTransactionRequest) (*pb.TransactionListResponse, error) {
	txs, err := h.App.GetAll(ctx, req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get all transactions: %v", err)
	}

	var pbTxs []*pb.Transaction
	for _, tx := range txs {
		pbTxs = append(pbTxs, mapToPbTransaction(tx))
	}

	return &pb.TransactionListResponse{
		Transactions: pbTxs,
	}, nil
}

// ✅ GetByIdTransaction
func (h *TransactionHandler) GetByIdTransaction(ctx context.Context, req *pb.GetByIdTransactionRequest) (*pb.TransactionResponse, error) {
	tx, err := h.App.GetById(ctx, req.GetId(), req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get transaction by id: %v", err)
	}
	if tx == nil {
		return nil, status.Errorf(codes.NotFound, "transaction not found")
	}

	return &pb.TransactionResponse{
		Transaction: mapToPbTransaction(tx),
	}, nil
}

// ✅ UpdateTransaction
func (h *TransactionHandler) UpdateTransaction(ctx context.Context, req *pb.UpdateTransactionRequest) (*pb.TransactionResponse, error) {
	tx := &domain.Transaction{
		ID:         req.GetId(),
		UserID:     req.GetUserId(),
		CourierID:  req.GetCourierId(),
		MerchantID: req.GetMerchantId(),
		Total:      req.GetTotal(),
		Status:     req.GetStatus(),
		SnapToken:  req.GetSnapToken(),
	}

	updatedTx, err := h.App.Update(ctx, tx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update transaction: %v", err)
	}

	return &pb.TransactionResponse{
		Transaction: mapToPbTransaction(updatedTx),
	}, nil
}

// ✅ DeleteTransaction
func (h *TransactionHandler) DeleteTransaction(ctx context.Context, req *pb.DeleteTransactionRequest) (*pb.DeleteTransactionResponse, error) {
	err := h.App.Delete(ctx, req.GetId(), req.GetUserId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete transaction: %v", err)
	}

	return &pb.DeleteTransactionResponse{
		Message: "Transaction deleted successfully",
	}, nil
}

func mapToPbTransaction(tx *domain.Transaction) *pb.Transaction {
	var pbDetails []*pb.TransactionDetail
	for _, d := range tx.Details {
		pbDetails = append(pbDetails, &pb.TransactionDetail{
			Id:            d.ID,
			TransactionId: d.TransactionID,
			FoodId:        d.FoodID,
			MerchantId:    d.MerchantID,
			Qty:           d.Qty,
			Price:         d.Price,
		})
	}

	return &pb.Transaction{
		Id:         tx.ID,
		UserId:     tx.UserID,
		CourierId:  tx.CourierID,
		MerchantId: tx.MerchantID,
		Total:      tx.Total,
		Status:     tx.Status,
		SnapToken:  tx.SnapToken,
		CreatedAt:  timestamppb.New(tx.CreatedAt),
		UpdatedAt:  timestamppb.New(tx.UpdatedAt),
		Details:    pbDetails, // ✅ penting!
	}
}
