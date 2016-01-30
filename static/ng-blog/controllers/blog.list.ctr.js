angular.module("blogApp")
    .controller("ArticleListCtrl", ["$scope", "BASE_CONFIG", "$stateParams", "articles", function ($scope, BASE_CONFIG, $stateParams, articles) {
        $scope.articles = articles.articles
        $scope.total_count = articles.total_count

        if ($stateParams.page >= 1) {
            $scope.perv_num = $stateParams.page - 1
        }

        if (articles.total_count / BASE_CONFIG.blogListPerPage > $stateParams.page) {
            $scope.next_num = $stateParams.page + 1
        }

        $scope.category = $stateParams.category
    }]);
