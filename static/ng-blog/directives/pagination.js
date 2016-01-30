angular.module("blogApp")
    .directive("myPagination", 
        function () {
            return {
                restrict: "E",
                templateUrl: "templates/pagination.html"
            };
        }
    );
