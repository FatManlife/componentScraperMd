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
        navigate(`?${newParams.toString()}`);
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
                        className="px-3 py-2 text-gray-500"
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
                    className={`px-3 py-2 rounded transition ${
                        page === currentPage
                            ? "bg-blue-500 text-white font-semibold"
                            : "bg-white text-gray-700 hover:bg-gray-100"
                    } disabled:cursor-not-allowed`}
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
                className="px-4 py-2 bg-white text-gray-700 rounded hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed transition"
            >
                Previous
            </button>

            {renderPageNumbers()}

            <button
                onClick={() => handlePageChange(currentPage + 1)}
                disabled={currentPage === totalPages}
                className="px-4 py-2 bg-white text-gray-700 rounded hover:bg-gray-100 disabled:opacity-50 disabled:cursor-not-allowed transition"
            >
                Next
            </button>
        </div>
    );
}

export default Pagination;
