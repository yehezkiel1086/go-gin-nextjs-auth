import { Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle } from "@/components/ui/card"
import { Badge } from "./ui/badge"
import { Button } from "@/components/ui/button"
import { MapPin, Building, Calendar, DollarSign } from "lucide-react"
import { Job } from "@/lib/api"

interface JobCardProps {
  job: Job;
}

export function JobCard({ job }: JobCardProps) {
  const formatDate = (dateString: string) => {
    return new Date(dateString).toLocaleDateString('en-US', {
      year: 'numeric',
      month: 'short',
      day: 'numeric'
    });
  };

  const formatSalary = (salary: number) => {
    return new Intl.NumberFormat('id-ID', {
      style: 'currency',
      currency: 'IDR',
      minimumFractionDigits: 0,
      maximumFractionDigits: 0,
    }).format(salary);
  };

  return (
    <Card className="w-full hover:shadow-lg transition-shadow duration-200">
      <CardHeader>
        <div className="flex justify-between items-start">
          <div className="flex-1">
            <CardTitle className="text-xl font-bold text-gray-900 mb-2">
              {job.title}
            </CardTitle>
            <div className="flex items-center gap-4 text-sm text-gray-600">
              <div className="flex items-center gap-1">
                <Building className="w-4 h-4" />
                <span>{job.company}</span>
              </div>
              <div className="flex items-center gap-1">
                <MapPin className="w-4 h-4" />
                <span>{job.location}</span>
              </div>
            </div>
          </div>
          <div className="text-right">
            <Badge variant="secondary" className="text-sm font-medium">
              {formatSalary(job.salary)}
            </Badge>
          </div>
        </div>
      </CardHeader>
      
      <CardContent>
        <CardDescription className="text-gray-700 line-clamp-3">
          {job.description}
        </CardDescription>
      </CardContent>
      
      <CardFooter className="flex justify-between items-center pt-4">
        <div className="flex items-center gap-1 text-xs text-gray-500">
          <Calendar className="w-3 h-3" />
          <span>Posted {formatDate(job.CreatedAt)}</span>
        </div>
        <Button size="sm" className="px-4">
          Apply Now
        </Button>
      </CardFooter>
    </Card>
  )
}
