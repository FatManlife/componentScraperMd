import { useNavigate, useSearchParams } from "react-router-dom";

type PaginationProps = {
    currentPage: number;
    totalCount: number;
    itemsPerPage?: number;
};

function Pagination({
    currentPage,
    totalCount,
    itemsPerPage = 24,
}: PaginationProps) {
    const navigate = useNavigate();
    const [searchParams] = useSearchParams();
    const totalPages = Math.ceil(totalCount / itemsPerPage);

    const handlePageChange = (page: number) => {
        const newParams = new URLSearchParams(searchParams);
        newParams.set("page", page.toString());
        navigate(`?${newParams.toString()}`, { replace: true });
    };

    if (totalPages <= 1) return null;

    const renderPageNumbers = () => {
        const pages = [];
        const maxVisible = 7;

        if (totalPages <= maxVisible) {
            for (let i = 1; i <= totalPages; i++) {
                pages.push(i);
            }
        } else {
            if (currentPage <= 4) {
                for (let i = 1; i <= 5; i++) {
                    pages.push(i);
                }
                pages.push("...");
                pages.push(totalPages);
            } else if (currentPage >= totalPages - 3) {
                pages.push(1);
                pages.push("...");
                for (let i = totalPages - 4; i <= totalPages; i++) {
                    pages.push(i);
                }
            } else {
                pages.push(1);
                pages.push("...");
                for (let i = currentPage - 1; i <= currentPage + 1; i++) {
                    pages.push(i);
                }
                pages.push("...");
                pages.push(totalPages);
            }
        }

        return pages.map((page, index) => {
            if (page === "...") {
                return (
                    <span
                        key={`ellipsis-${index}`}
                        className="px-3 py-2"
                        style={{ color: '#8A8A8A' }}
                    >
                        ...
                    </span>
                );
            }

            return (
                <button
                    key={page}
                    onClick={() => handlePageChange(page as number)}
                    disabled={page === currentPage}
                    className="px-3 py-2 transition-all disabled:cursor-not-allowed"
                    style={{
                        backgroundColor: page === currentPage ? '#000000' : '#FFFFFF',
                        color: page === currentPage ? '#FFFFFF' : '#000000',
                        border: '1px solid #D9D9D9',
                        borderRadius: '2px',
                        fontWeight: page === currentPage ? '600' : '400'
                    }}
                    onMouseEnter={(e) => {
                        if (page !== currentPage) {
                            e.currentTarget.style.backgroundColor = '#F4F4F4';
                        }
                    }}
                    onMouseLeave={(e) => {
                        if (page !== currentPage) {
                            e.currentTarget.style.backgroundColor = '#FFFFFF';
                        }
                    }}
                >
                    {page}
                </button>
            );
        });
    };

    return (
        <div className="flex items-center justify-center gap-2 mt-6">
            <button
                onClick={() => handlePageChange(currentPage - 1)}
                disabled={currentPage === 1}
                className="px-4 py-2 transition-all disabled:cursor-not-allowed"
                style={{
                    backgroundColor: '#FFFFFF',
                    color: '#000000',
                    border: '1px solid #D9D9D9',
                    borderRadius: '2px',
                    opacity: currentPage === 1 ? 0.5 : 1
                }}
                onMouseEnter={(e) => {
                    if (currentPage !== 1) {
                        e.currentTarget.style.backgroundColor = '#F4F4F4';
                    }
                }}
                onMouseLeave={(e) => {
                    if (currentPage !== 1) {
                        e.currentTarget.style.backgroundColor = '#FFFFFF';
                    }
                }}
            >
                Previous
            </button>

            {renderPageNumbers()}

            <button
                onClick={() => handlePageChange(currentPage + 1)}
                disabled={currentPage === totalPages}
                className="px-4 py-2 transition-all disabled:cursor-not-allowed"
                style={{
                    backgroundColor: '#FFFFFF',
                    color: '#000000',
                    border: '1px solid #D9D9D9',
                    borderRadius: '2px',
                    opacity: currentPage === totalPages ? 0.5 : 1
                }}
                onMouseEnter={(e) => {
                    if (currentPage !== totalPages) {
                        e.currentTarget.style.backgroundColor = '#F4F4F4';
                    }
                }}
                onMouseLeave={(e) => {
                    if (currentPage !== totalPages) {
                        e.currentTarget.style.backgroundColor = '#FFFFFF';
                    }
                }}
            >
                Next
            </button>
        </div>
    );
}

export default Pagination;
