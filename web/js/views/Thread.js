export default async function ThreadView() {
    return `<div class="focus-post">
                <h2>@mohamed</h2>

                <div class="frame">
                    <h3>clickbait</h3>
                    <div class="content">
                        <pre>clickbaitrt</pre>
                    </div>
                    <div class="post-meta" style="display: flex; justify-content: space-between;">

                        <form style="margin-top: 10px;" method="POST" action="/like" style="display: inline;">
                            <input type="hidden" name="post_id" value="{{.Post.ID}}">
                            <button type="submit" name="type" value="like" style="display: inline;"
                                class="like-btn">
                                <span class="material-icons">thumb_up</span>
                                <span>22</span>
                            </button>
                            <button type="submit" name="type" value="dislike" style="display: inline;"
                                class="dislike-btn">
                                <span class="material-icons">thumb_down</span>
                                <span>6</span>
                            </button>
                        </form>

                        <div class="category-flairs" style="display: inline; margin-top: 10px;">
                            <span class="category-flair">music</span>
                            <span class="category-flair">gaming</span>
                        </div>
                    </div>


                </div>
            </div>

            <div class="comment-input-container">
                <div class="comment-user">
                    <div class="username">
                        <h2>@michaeljordan</h2>
                    </div>
                </div>
                <div class="comment-input-section" id="comment-input">
                    <form method="POST" action="/comment">
                        <input type="hidden" name="post_id" value="{{.Post.ID}}">
                        <textarea name="content" id="comment-input" placeholder="Write your comment..."></textarea>
                        <button class="post-comment-btn">Post Comment</button>
                    </form>
                </div>
            </div>

            <!-- place a comment count here -->
        <div class="comment-count">
                <span>23 Comments</span>
        </div>
        <div class="comment-block" style="display: block;">
            <div class="comment-user">
                <div class="username">
                    <h2>james</h2>
                </div>
            </div>
            <div class="content1">
                <pre>sdnsdjnsjkdsk</pre>
            </div>

            <div class="comment-meta">
                <form method="POST" action="/like">
                    <input type="hidden" name="post_id" value="{{.PostID}}">
                    <input type="hidden" name="comment_id" value="{{.ID}}">

                    <button type="submit" name="type" value="like" style="display: inline;"
                        class="like-btn">
                        <span class="material-icons">thumb_up</span>
                        <span>22</span>
                    </button>
                    <button type="submit" name="type" value="dislike" style="display: inline;"
                        class="dislike-btn">
                        <span class="material-icons">thumb_down</span>
                        <span>2</span>
                    </button>
                </form>
            </div>
        </div>
        <p style="color: white;">No Comments yet, be the first to comment!</p>
    </div>
`;
}

