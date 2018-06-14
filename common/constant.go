package common

import "regexp"

const (
	LOGIN_SESSION_KEY = "login_user"
	USER_IN_COOKIE = "S_L_ID"
	/**
     * aes加密加盐
     */
	AES_SALT = "0123456789abcdef"
	/**
     * 最大获取文章条数
     */
	MAX_POSTS = 9999

	/**
     * 最大页码
     */
	MAX_PAGE = 100

	/**
     * 文章最多可以输入的文字数
     */
	MAX_TEXT_COUNT = 200000

	/**
     * 文章标题最多可以输入的文字个数
     */
	MAX_TITLE_COUNT = 200

	/**
     * 点击次数超过多少更新到数据库
     */
	HIT_EXCEED = 10

	/**
     * 上传文件最大1M
     */
	MAX_FILE_SIZE = 1048576

	/**
     * 成功返回
     */
	SUCCESS_RESULT  = "SUCCESS"

	/**
     * 同一篇文章在2个小时内无论点击多少次只算一次阅读
     */
	HITS_LIMIT_TIME = 7200


	SLUG_REGEX_PATTERN = "^[A-Za-z0-9_-]{5,100}$"

	VALID_EMAIL_ADDRESS_REGEX_PATTERN = `^[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,6}$`

)


var (
	SLUG_REGEX,_ = regexp.Compile(SLUG_REGEX_PATTERN)
	VALID_EMAIL_ADDRESS_REGEX,_ = regexp.Compile(VALID_EMAIL_ADDRESS_REGEX_PATTERN)
)