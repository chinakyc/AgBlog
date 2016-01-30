angular.module('blogApp')
  .filter('htmlToPlaintext', function() {
            return function(text) {
                return  text ? String(text).replace(/<[^>]+>/gm, '').slice(0, 300) : '';
            };
        }
    );
