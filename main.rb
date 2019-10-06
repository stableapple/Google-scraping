require 'nokogiri'
require 'open-uri'
doc= Nokogiri::HTML(open('https://www.google.com/search?client=ubuntu&channel=fs&q=linux&ie=utf-8&oe=utf-8'))
doc.css('cite.iUh30').each do |link|
puts link.content
end







