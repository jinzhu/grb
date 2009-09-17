module Git
  def push_to(from,to)
    '"#{GIT} push #{origin} #{current_branch}:refs/heads/#{branch}"',
    fetch_and_rebase && push_to(from,to) if fail
  end

  def fetch_and_rebase
    '"#{GIT} fetch #{origin}"'
    '"#{GIT} rebase (related)"'
  end

  def track(from,to)
    '"#{GIT} branch --track #{branch} #{origin}/#{branch}"',
  end

  def add_config
    '"#{GIT} config branch.#{branch}.remote #{origin}"',
    '"#{GIT} config branch.#{branch}.merge refs/heads/#{branch}"',
  end
end
