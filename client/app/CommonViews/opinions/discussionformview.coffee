class DiscussionFormView extends KDFormView

  constructor :(options, data)->

    super

    {profile} = KD.whoami()

    @submitDiscussionBtn = new KDButtonView
      title           : "Save your changes"
      type            : "submit"
      cssClass        : "clean-gray discussion-submit-button"

    @discussionBody = new KDInputView
      cssClass        : "discussion-body"
      name            : "body"
      title           : "your Discussion Topic"
      type            : "textarea"
      # autogrow        : yes
      placeholder     : "What do you want to contribute to the discussion?"

    @discussionTitle = new KDInputView
      cssClass        : "discussion-title"
      name            : "title"
      title           : "your Opinion"
      type            : "text"
      # autogrow        : yes
      placeholder     : "What do you want to talk about?"

    @labelAddTags = new KDLabelView
      title           : "Add Tags:"

    @selectedItemWrapper = new KDCustomHTMLView
      tagName         : "div"
      cssClass        : "tags-selected-item-wrapper clearfix"

    @markdownLink = new KDCustomHTMLView
      tagName     : 'a'
      name        : "markdownLink"
      value       : "markdown is enabled"
      attributes  :
        title     : "markdown is enabled"
        href      : '#'
        value     : "markdown syntax is enabled"
      cssClass    : 'markdown-link'
      partial     : "markdown is enabled<span></span>"
      click       : (pubInst, event)=>
        if $(event.target).is 'span'
          link.hide()
        else
          markdownText = new KDMarkdownModalText
          modal = new KDModalView
            title       : "How to use the <em>markdown</em> syntax."
            cssClass    : "what-you-should-know-modal markdown-cheatsheet"
            height      : "auto"
            width       : 500
            content     : markdownText.markdownText()
            buttons     :
              Close     :
                title   : 'Close'
                style   : 'modal-clean-gray'
                callback: -> modal.destroy()

    @markdownSelect = new KDSelectBox
      type          : "select"
      name          : "markdown"
      cssClass      : "select markdown-select hidden"
      selectOptions :
          [
              title     : "enable markdown syntax"
              value     : "markdown"
            ,
              title     : "disable markdown syntax"
              value     : "nomarkdown"
          ]
      defaultValue  : "markdown"
      callback      : (value) =>
        @emit "opinion.changeMarkdown", value

    if data instanceof bongo.api.JDiscussion
      @discussionBody.setValue Encoder.htmlDecode data.body
      @discussionTitle.setValue Encoder.htmlDecode data.title

    @on "discussion.changeMarkdown", (value) ->

    # TODO keep tags on editing -arvid
    @tagController = new TagAutoCompleteController
      name                : "meta.tags"
      type                : "tags"
      itemClass           : TagAutoCompleteItemView
      selectedItemClass   : TagAutoCompletedItemView
      outputWrapper       : @selectedItemWrapper
      selectedItemsLimit  : 5
      listWrapperCssClass : "tags"
      itemDataPath        : 'title'
      form                : @
      dataSource          : (args, callback)=>
        {inputValue} = args
        updateWidget = @getDelegate()
        blacklist = (data.getId() for data in @tagController.getSelectedItemData() when 'function' is typeof data.getId)
        appManager.tell "Topics", "fetchTopics", {inputValue, blacklist}, callback

    @tagAutoComplete = @tagController.getView()

  viewAppended:()->

    @setClass "update-options discussion"
    @setTemplate @pistachio()
    @template.update()

  submit:=>
    @once "FormValidationPassed", => @reset()
    super


  pistachio:->
      """
      <div class="discussion-box">
        <div class="discussion-form">
          {{> @discussionTitle}}
          {{> @discussionBody}}
        </div>
        <div class="discussion-buttons">
          <div class="discussion-submit">
            {{> @markdownLink}}
            {{> @submitDiscussionBtn}}
          </div>
        </div>
      </div>
      """