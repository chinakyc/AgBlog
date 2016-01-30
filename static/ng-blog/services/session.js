// https://medium.com/opinionated-angularjs/techniques-for-authentication-in-angularjs-applications-7bbf0346acec#.nno3rovhb
// session to save userinfo
angular.module("blogApp").service("Session", ["$rootScope", "USER_ROLES", "MESSAGE_TYPE", "MESSAGE_EVEN",
    function($rootScope, USER_ROLES, MESSAGE_TYPE, MESSAGE_EVEN) {
        // save flashed messages
        this.msgBox = []
        this.userToken = null;
        this.userRole = null;
        this.user = null;
        // use for go back
        this.previousState = null;

        this.setPreviousState = function(state, params) {
            this.previousState = {state: state, params: params}
        };

        this.flashMessage = function (messageContent, messageType) {
            if (!messageType) {
                messageType = MESSAGE_TYPE.success;
            }

            this.msgBox.push({
                content: messageContent,
                type: messageType,
            });
            // fire message flashed even
            $rootScope.$broadcast(MESSAGE_EVEN.flashed);
        };

        this.getMessages = function (messageContent, messageType) {
            var msgBox = this.msgBox;
            this.msgBox = [];
            return msgBox
        };

        this.createUser = function(userToken, userRole, user) {
            this.userToken = userToken;
            this.userRole = userRole;
            this.user = user;
        };

        // destroy data in session
        this.destroy = function() {
            this.userToken = null;
            this.userRole = null;
            this.user = null;
            this.msgBox = [];
        };


        return this
    }
]);
