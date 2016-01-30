angular.module("blogApp").service("CategoryService", ["$http", "API_URLS",
    function($http, API_URLS) {

        // get all categorys
        this.get = function() {
            return $http.get(API_URLS.categorys);
        };

        return this;
    }
]);
