class Grb
  GIT    = ENV['GRB_GIT']    || 'git'
  ORIGIN = ENV['GRB_ORIGIN'] || 'origin'

  COMMANDS = {
    :new     => {
      :desc  => "=> create new branch `branch`\ngrb new [branch] [--explain]",
      :commands => [
        '"#{GIT} push #{origin} #{current_branch}:refs/heads/#{branch}"',
        '"#{GIT} fetch #{origin}"',
        '"#{GIT} branch --track #{branch} #{origin}/#{branch}"',
        '"#{GIT} checkout #{branch}"'
      ]
    },

    :push     => {
      :desc  => "=> push branch `branch`, default current_branch\ngrb push [branch] [--explain]",
      :commands => [
        '"#{GIT} push #{origin} #{branch}:refs/heads/#{branch}"',
        '"#{GIT} fetch #{origin}"',
        '"#{GIT} config branch.#{branch}.remote #{origin}"',
        '"#{GIT} config branch.#{branch}.merge refs/heads/#{branch}"',
        '"#{GIT} checkout #{branch}"'
      ]
    },

    :mv     => {
      :desc  => "=> rename `branch1` to `branch2`\ngrb mv   [branch1] [branch2] [--explain]\n=> rename current branch to `branch`\ngrb mv branch [--explain]",
      :commands => [
        ' if(branch != branch_)
           "#{GIT} push #{origin} #{branch}:refs/heads/#{branch_}
            #{GIT} fetch #{origin}
            #{GIT} branch --track #{branch_} #{origin}/#{branch_}
            #{GIT} checkout #{branch_}
            #{GIT} branch -d #{branch}
            #{GIT} push #{origin} :refs/heads/#{branch}"
          else
           "#{GIT} push #{origin} #{current_branch}:refs/heads/#{branch}
            #{GIT} fetch #{origin}
            #{GIT} branch --track #{branch} #{origin}/#{branch}
            #{GIT} checkout #{branch}
            #{GIT} push #{origin} :refs/heads/#{current_branch}
            #{GIT} branch -d #{current_branch}"
          end'
      ]
    },

    :rm     => {
      :desc  => "=> delete branch `branch`,default current_branch\ngrb rm [branch] [--explain]",
      :commands => [
        '"#{GIT} push #{origin} :refs/heads/#{branch}"',
        '"#{GIT} checkout master" if current_branch == branch',
        '"#{GIT} branch -d #{branch}"'
      ]
    },

    :pull      => {
      :desc  => "=> pull branch `branch`,default current_branch\ngrb pull [branch] [--explain]",
      :commands => [
        '"#{GIT} fetch #{origin}"',
        'if local_branches.include?(branch) 
          "#{GIT} config branch.#{branch}.remote #{origin}\n" +
          "#{GIT} config branch.#{branch}.merge refs/heads/#{branch}"
        else
          "#{GIT} branch --track #{branch} #{origin}/#{branch}"
        end'
      ]
    }
  }

  def self.parse(opt)
    if COMMANDS.keys.include?(opt[:command].to_sym)
      current_branch,branch,branch_,origin = get_current_branch,opt[:branch],opt[:branch_],ORIGIN

      COMMANDS[opt[:command].to_sym][:commands].map {|x| exec_cmd(eval(x))}
    else
     help 
    end
  end

  def self.exec_cmd(str)
    return true unless str
    puts("\e[031m" + str.gsub(/^\s*/,'') + "\e[0m")
    system("#{str}") unless EXPLAIN
  end

  def self.get_current_branch
    (`git branch 2> /dev/null | grep '^\*'`).gsub(/\W/,'')
  end

  def self.local_branches
   (`git branch -l`).split(/\n/).map{|x| x.gsub(/\W/,'')}
  end

  def self.help(*args)
    puts "USAGE:"
    COMMANDS.values.map {|x| puts x[:desc].gsub(/^(\W.*)$/,"\e[31m" + '\1' + "\e[0m").gsub(/^(\w.*)$/,'  $ \1')}
  end
end
